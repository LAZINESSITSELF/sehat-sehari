import numpy as np

# Hitung BMR menggunakan Mifflin-St Jeor
def calculate_bmr(age, gender, height_cm, weight_kg):
    if gender == "male":
        return 10 * weight_kg + 6.25 * height_cm - 5 * age + 5
    else:
        return 10 * weight_kg + 6.25 * height_cm - 5 * age - 161

# Faktor aktivitas (TDEE multiplier)
ACTIVITY_MULTIPLIER = {
    "sedentary": 1.2,
    "light": 1.375,
    "moderate": 1.55,
    "active": 1.725,
    "very_active": 1.9
}

# Sesuaikan berdasarkan goal
GOAL_ADJUSTMENT = {
    "lose_weight": -500,     # defisit 500 kalori
    "maintain_weight": 0,
    "gain_weight": 500       # surplus 500 kalori
}

# Hitung TDEE dan calorie_need
def calculate_calorie_need(profile):
    bmr = calculate_bmr(profile.age, profile.gender, profile.height_cm, profile.weight_kg)
    tdee = bmr * ACTIVITY_MULTIPLIER[profile.activity_level]
    adjustment = GOAL_ADJUSTMENT[profile.goal]
    return round(tdee + adjustment, 2)

# Rekomendasi makronutrien (asumsi: 
#    protein 20%, fat 25%, carbs 55% dari total kalori)
def calculate_macros(calorie_need):
    # 1 gram protein = 4 kcal; 1 gram lipida = 9 kcal; 1 gram karbohidrat = 4 kcal
    protein_kcal = calorie_need * 0.20
    fat_kcal = calorie_need * 0.25
    carbs_kcal = calorie_need * 0.55

    return {
        "protein_g": round(protein_kcal / 4, 1),
        "fat_g": round(fat_kcal / 9, 1),
        "carbs_g": round(carbs_kcal / 4, 1)
    }

# Contoh dataset statis (boleh diganti dinamis dari file/data nyata)
SAMPLE_MENU = [
    {"name": "Oatmeal dengan buah", "calories": 350, "macros": {"protein_g": 10, "fat_g": 5, "carbs_g": 60}},
    {"name": "Salad ayam panggang", "calories": 450, "macros": {"protein_g": 30, "fat_g": 15, "carbs_g": 40}},
    {"name": "Nasi merah + ikan bakar", "calories": 500, "macros": {"protein_g": 25, "fat_g": 10, "carbs_g": 65}},
    {"name": "Smoothie hijau", "calories": 250, "macros": {"protein_g": 5, "fat_g": 3, "carbs_g": 45}},
]

# Contoh dataset aktivitas
SAMPLE_ACTIVITIES = [
    {"name": "Jalan cepat", "duration_min": 30, "calories_burned": 150},
    {"name": "Bersepeda ringan", "duration_min": 30, "calories_burned": 200},
    {"name": "Yoga", "duration_min": 45, "calories_burned": 180},
    {"name": "Lari ringan", "duration_min": 30, "calories_burned": 300},
]

def recommend_menu(calorie_need):
    # Pilih menu yang total kalorinya mendekati 70%-80% dari calorie_need
    target = calorie_need * 0.75
    # Sederhana: random sampling beberapa menu hingga mendekati target
    menu = []
    total = 0
    for item in SAMPLE_MENU:
        if total + item["calories"] <= target:
            menu.append(item)
            total += item["calories"]
    return menu

def recommend_activities(profile):
    # Pilih aktivitas sesuai durasi harian
    activities = []
    remaining = profile.exercise_duration_min
    for item in SAMPLE_ACTIVITIES:
        if remaining <= 0:
            break
        if item["duration_min"] <= remaining:
            activities.append(item)
            remaining -= item["duration_min"]
    return activities
