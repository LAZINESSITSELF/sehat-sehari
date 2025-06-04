from flask import Flask, request, jsonify
import numpy as np
from sklearn.ensemble import RandomForestRegressor
import os

app = Flask(__name__)

def preprocess_input(data):
    """
    Mengubah input terstruktur menjadi array 1×6 untuk prediksi.
    """
    gender = 0 if data.get("gender", "male").lower() == "male" else 1
    al = data.get("activity_level", "sedentary").lower()
    if al == "moderate":
        activity_factor = 1.55
    elif al == "active":
        activity_factor = 1.9
    else:
        activity_factor = 1.2
    exercise = float(data.get("exercise_minutes", 0))
    return np.array([[\
        float(data.get("age", 30)),
        float(data.get("height", 170)),
        float(data.get("weight", 70)),
        gender,
        activity_factor,
        exercise
    ]])

def generate_training_data(n_samples=200):
    """
    Membuat data sintetis untuk melatih RandomForestRegressor.
    Fitur: [age, height, weight, gender, activity_factor, exercise_minutes]
    Target: total calories (BMR * activity_factor + exercise*3 + noise)
    """
    np.random.seed(42)
    X, y = [], []
    for _ in range(n_samples):
        age = np.random.uniform(18, 65)
        height = np.random.uniform(150, 200)
        weight = np.random.uniform(50, 100)
        gender = np.random.choice([0, 1])  # 0: male, 1: female
        activity_factor = np.random.choice([1.2, 1.55, 1.9])
        exercise = np.random.uniform(0, 120)
        if gender == 0:
            bmr = 10 * weight + 6.25 * height - 5 * age + 5
        else:
            bmr = 10 * weight + 6.25 * height - 5 * age - 161
        calories = bmr * activity_factor + 3 * exercise
        noise = np.random.normal(0, 50)
        calories += noise
        X.append([age, height, weight, gender, activity_factor, exercise])
        y.append(calories)
    return np.array(X), np.array(y)

# Latih model sekali saat startup
X_train, y_train = generate_training_data(200)
model = RandomForestRegressor(n_estimators=100, random_state=42)
model.fit(X_train, y_train)

def calculate_recommendation_advanced(data):
    """
    Menghasilkan rekomendasi kalori, menu, dan olahraga dengan opsi variatif.
    data (dict) must contain keys:
      - age, height, weight, gender, activity_level, exercise_minutes, goal
    """
    try:
        X_new = preprocess_input(data)
    except Exception as e:
        return None, "Input data tidak valid: " + str(e)
    
    # Prediksi kalori dasar
    base_calories = model.predict(X_new)[0]
    goal = data.get("goal", "maintain").lower()
    if goal == "lose":
        predicted_calories = base_calories * 0.85
    elif goal == "gain":
        predicted_calories = base_calories * 1.15
    else:
        predicted_calories = base_calories

    # Pilih rekomendasi menu sesuai goal
    if goal == "lose":
        menu_recommendation = "Salad sayur + dada ayam panggang + oatmeal"
        menu_options = [
            "Salad sayur + dada ayam panggang + oatmeal",
            "Sayur kukus + ikan kukus + quinoa",
            "Oatmeal + telur rebus + buah apel"
        ]
    elif goal == "gain":
        menu_recommendation = "Nasi merah + ikan salmon panggang + alpukat"
        menu_options = [
            "Nasi merah + ikan salmon panggang + alpukat",
            "Steak daging tanpa lemak + kentang rebus + brokoli",
            "Smoothie pisang + selai kacang + susu almond"
        ]
    else:  # maintain
        menu_recommendation = "Kombinasi sayur, protein tanpa lemak, dan karbohidrat kompleks"
        menu_options = [
            "Kombinasi sayur, protein tanpa lemak, dan karbohidrat kompleks",
            "Nasi merah + dada ayam panggang + sayur tumis",
            "Oatmeal + Greek yogurt + buah berries"
        ]

    # Pilih rekomendasi olahraga sesuai activity_level + goal
    al = data.get("activity_level", "sedentary").lower()
    if al == "active":
        if goal == "lose":
            sport_recommendation = "HIIT 40 menit atau berenang"
            sport_options = ["HIIT 40 menit", "Berenang 30 menit", "Lari 45 menit"]
        elif goal == "gain":
            sport_recommendation = "Angkat beban + lari ringan 20 menit"
            sport_options = ["Angkat beban 45 menit", "Lari ringan 20 menit", "Senam kekuatan 30 menit"]
        else:  # maintain
            sport_recommendation = "Lari santai 30 menit atau yoga"
            sport_options = ["Lari santai 30 menit", "Yoga 40 menit", "Senam kardio ringan 30 menit"]
    elif al == "moderate":
        if goal == "lose":
            sport_recommendation = "Jogging 30 menit + senam ringan"
            sport_options = ["Jogging 30 menit", "Senam ringan 30 menit", "Bersepeda santai 30 menit"]
        elif goal == "gain":
            sport_recommendation = "Sepeda santai 45 menit"
            sport_options = ["Sepeda santai 45 menit", "Angkat beban ringan 30 menit", "Latihan calisthenics 30 menit"]
        else:  # maintain
            sport_recommendation = "Jogging ringan 30 menit"
            sport_options = ["Jogging ringan 30 menit", "Senam ringan 30 menit", "Berjalan cepat 30 menit"]
    else:  # sedentary
        if goal == "lose":
            sport_recommendation = "Jalan kaki 30 menit + peregangan"
            sport_options = ["Jalan kaki 30 menit", "Peregangan 20 menit", "Senam ringan 20 menit"]
        elif goal == "gain":
            sport_recommendation = "Yoga ringan + naik turun tangga"
            sport_options = ["Yoga ringan 30 menit", "Naik turun tangga 20 menit", "Senam ringan 20 menit"]
        else:  # maintain
            sport_recommendation = "Jalan santai 20 menit"
            sport_options = ["Jalan santai 20 menit", "Peregangan 15 menit", "Senam ringan 20 menit"]

    result = {
        "calories_needed": int(round(predicted_calories)),
        "menu_recommendation": menu_recommendation,
        "menu_options": menu_options,
        "sport_recommendation": sport_recommendation,
        "sport_options": sport_options
    }
    return result, None

@app.route('/predict', methods=['POST'])
def predict():
    data = request.get_json()
    if not data:
        return jsonify({"error": "Tidak ada data input."}), 400
    recommendation, error = calculate_recommendation_advanced(data)
    if error:
        return jsonify({"error": error}), 400
    return jsonify(recommendation), 200

if __name__ == '__main__':
    port = int(os.getenv("ML_PORT", 5000))
    app.run(host='0.0.0.0', port=port, debug=True)