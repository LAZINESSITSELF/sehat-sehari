from pydantic import BaseModel
from typing import Literal, List, Optional

class UserProfile(BaseModel):
    age: int
    gender: Literal["male", "female"]
    height_cm: float
    weight_kg: float
    activity_level: Literal["sedentary", "light", "moderate", "active", "very_active"]
    goal: Literal["lose_weight", "maintain_weight", "gain_weight"]
    exercise_duration_min: int  # rata-rata menit per hari

class Macros(BaseModel):
    protein_g: float
    fat_g: float
    carbs_g: float

class MenuItem(BaseModel):
    name: str
    calories: float
    macros: Macros

class ActivityItem(BaseModel):
    name: str
    duration_min: int
    calories_burned: float

class Recommendation(BaseModel):
    calorie_need: float
    macros: Macros
    menu: List[MenuItem]
    activities: List[ActivityItem]
