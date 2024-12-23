import { createStore } from 'vuex';

export default createStore({
    state: {
        token: null,
        role: null,
    },
    mutations: {
        setToken(state, token) {
        state.token = token;
        },
        setRole(state, role) {
        state.role = role;
        },
    },
    actions: {
        login({ commit }, { token, role }) {
        commit('setToken', token);
        commit('setRole', role);
        },
        logout({ commit }) {
        commit('setToken', null);
        commit('setRole', null);
        },
    },
    getters: {
        isAuthenticated: (state) => !!state.token,
        getRole: (state) => state.role,
    },
});
