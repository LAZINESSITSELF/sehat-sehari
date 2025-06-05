import axios from "axios";

const API_BASE = "http://localhost:8080/api"; // backend Golang

const api = axios.create({
    baseURL: API_BASE,
    headers: {
        "Content-Type": "application/json",
    },
});

// Interceptor untuk menyisipkan token Authorization jika ada
api.interceptors.request.use((config) => {
    const token = localStorage.getItem("token");
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

export default api;
