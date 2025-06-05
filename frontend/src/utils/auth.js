// src/utils/auth.js

/**
 * Hapus token dari localStorage dan arahkan pengguna ke halaman login.
 * @param {Router} router - instance Vue Router
 */
export function logout(router) {
    localStorage.removeItem("token");

    router.push("/login");
}
  