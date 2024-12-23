import axios from "axios";

// Konfigurasi default Axios
const axiosInstance = axios.create({
    baseURL: "http://localhost:8080", // Ganti dengan URL back-end Anda
    headers: {
        "Content-Type": "application/json",
    },
});

// Interceptor untuk menyisipkan token (jika diperlukan)
axiosInstance.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem("token"); // Ambil token dari localStorage
        if (token) {
        config.headers["Authorization"] = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

export default axiosInstance;
