# ml_service.py
from flask import Flask, request, jsonify
import numpy as np
from sklearn.ensemble import RandomForestRegressor
import os

app = Flask(__name__)

# --- Preprocessing & Data Simulation Functions ---

def preprocess_input(data):
    """
    Mengubah input JSON ke dalam format angka (fitur numerik) untuk prediksi.
    Input dijadikan array dengan urutan:
    [age, height, weight, gender, activity_factor, exercise_minutes]
    
    - gender: "male" → 0, "female" → 1
    - activity_level: "sedentary" → 1.2, "moderate" → 1.55, "active" → 1.9  
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
    return np.array([[
        float(data.get("age", 30)),
        float(data.get("height", 170)),
        float(data.get("weight", 70)),
        gender,
        activity_factor,
        exercise
    ]])

def generate_training_data(n_samples=200):
    """
    Menghasilkan dataset simulasi dengan n_samples.
    Fitur:
      - age (18-65 tahun)
      - height (150-200 cm)
      - weight (50-100 kg)
      - gender: 0 (male) atau 1 (female)
      - activity_factor: 1.2, 1.55, atau 1.9
      - exercise_minutes (0-120 menit)
    Target (calories) dihitung berdasarkan:
      - BMR Mifflin-St Jeor:
          * pria: (10×weight) + (6.25×height) – (5×age) + 5
          * wanita: (10×weight) + (6.25×height) – (5×age) – 161
      - Target = BMR × activity_factor + (3 × exercise_minutes) + noise  
    """
    np.random.seed(42)
    X = []
    y = []
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
        # Setiap menit olahraga menambah sekitar 3 kalori
        calories = bmr * activity_factor + 3 * exercise
        noise = np.random.normal(0, 50)
        calories += noise
        X.append([age, height, weight, gender, activity_factor, exercise])
        y.append(calories)
    return np.array(X), np.array(y)

# --- Training the Model ---

X_train, y_train = generate_training_data(200)
model = RandomForestRegressor(n_estimators=100, random_state=42)
model.fit(X_train, y_train)

# --- Advanced Prediction Function ---

def calculate_recommendation_advanced(data):
    try:
        X_new = preprocess_input(data)
    except Exception as e:
        return None, "Input data tidak valid: " + str(e)
    
    base_calories = model.predict(X_new)[0]
    goal = data.get("goal", "maintain").lower()
    if goal == "lose":
        predicted_calories = base_calories * 0.85
    elif goal == "gain":
        predicted_calories = base_calories * 1.15
    else:
        predicted_calories = base_calories

    al = data.get("activity_level", "sedentary").lower()
    if al == "moderate":
        sport_recommendation = "Jogging ringan 30 menit"
    elif al == "active":
        sport_recommendation = "Lari atau HIIT 30 menit"
    else:  # sedentary atau default
        sport_recommendation = "Jalan santai 20 menit"

    menu_recommendation = "Kombinasi sayur, protein tanpa lemak, dan karbohidrat kompleks"
    result = {
        "calories_needed": int(predicted_calories),
        "menu_recommendation": menu_recommendation,
        "sport_recommendation": sport_recommendation
    }
    return result, None

# --- Flask Endpoint ---

@app.route('/predict', methods=['POST'])
def predict():
    data = request.get_json()
    if not data:
        return jsonify({"error": "Tidak ada data input."}), 400
    recommendation, error = calculate_recommendation_advanced(data)
    if error:
        return jsonify({"error": error}), 400
    return jsonify(recommendation), 200

# --- Menjalankan Aplikasi Flask ---

if __name__ == '__main__':
    port = int(os.getenv("ML_PORT", 5000))
    app.run(host='0.0.0.0', port=port, debug=True)
