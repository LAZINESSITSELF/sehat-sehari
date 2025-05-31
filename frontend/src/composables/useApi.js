import api from '../api/api.js';

export async function loginUser(formData) {
    const response = await fetch(api.login, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData)
    })
    if (!response.ok) {
        throw new Error("Login failed")
    }
    return response.json()
}

export async function registerUser(formData) {
    const response = await fetch(api.register, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData)
    })
    if (!response.ok) {
        throw new Error("Registration failed")
    }
    return response.json()
}

export async function fetchRecommendation(formData) {
    const response = await fetch(api.recommendation, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData)
    })
    if (!response.ok) {
        throw new Error("Failed to fetch recommendation")
    }
    return response.json()
  }