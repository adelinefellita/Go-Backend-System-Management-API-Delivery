<template>
	<div>
		<h1>Login</h1>
		<form @submit.prevent="handleLogin">
			<input v-model="email" type="email" placeholder="Email" required />
			<input v-model="password" type="password" placeholder="Password" required />
			<button type="submit">Login</button>
		</form>
	</div>
</template>
<script>
	import axios from 'axios';
	export default {
		data() {
			return {
				email: '',
				password: '',
			};
		},
		methods: {
			async handleLogin() {
				try {
					const response = await axios.post('http://localhost:8080/login', {
						email: this.email,
						password: this.password,
					});
					const {
						token,
						role
					} = response.data;
					this.$store.dispatch('login', {
						token,
						role
					});
					if (role === 'manager') {
						this.$router.push('/manager');
					} else if (role === 'courier') {
						this.$router.push('/courier');
					}
				} catch (error) {
					alert('Login failed. Please check your credentials.');
				}
			},
		},
	};
</script>