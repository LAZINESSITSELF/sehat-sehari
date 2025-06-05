from fastapi import FastAPI
from models import UserProfile, Recommendation, MenuItem, ActivityItem, Macros
from utils import (
    calculate_calorie_need, calculate_macros,
    recommend_menu, recommend_activities
)

app = FastAPI(title="ML Service – Sehat Sehari")

@app.post("/api/recommend", response_model=Recommendation)
def get_recommendation(profile: UserProfile):
    # Hitung kebutuhan kalori
    calorie_need = calculate_calorie_need(profile)
    # Hitung makro
    macros_dict = calculate_macros(calorie_need)
    macros = Macros(**macros_dict)

    # Rekomendasi menu
    menu_raw = recommend_menu(calorie_need)
    menu = [
        MenuItem(
            name=item["name"],
            calories=item["calories"],
            macros=Macros(**item["macros"])
        )
        for item in menu_raw
    ]

    # Rekomendasi aktivitas
    activities_raw = recommend_activities(profile)
    activities = [
        ActivityItem(
            name=item["name"],
            duration_min=item["duration_min"],
            calories_burned=item["calories_burned"]
        )
        for item in activities_raw
    ]

    return Recommendation(
        calorie_need=calorie_need,
        macros=macros,
        menu=menu,
        activities=activities
    )
