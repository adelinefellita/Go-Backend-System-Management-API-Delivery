<template>
	<div>
		<h1>Manager Dashboard</h1>
		<form @submit.prevent="addAddress">
			<!-- Form input for adding address -->
		</form>
		<ul>
			<li v-for="address in addresses" :key="address.id">
				{{ address }}
				<button @click="deleteAddress(address.id)">Delete</button>
			</li>
		</ul>
	</div>
</template>
<script>
	import axios from 'axios';
	export default {
		data() {
			return {
				addresses: [],
			};
		},
		async created() {
			const response = await axios.get('/api/manager/addresses', {
				headers: {
					Authorization: `Bearer ${this.$store.state.token}`,
				},
			});
			this.addresses = response.data;
		},
		methods: {
			async addAddress() {
				// Implement add address
			},
			async deleteAddress(id) {
				await axios.delete(`http://localhost:8080/manager/addresses/${id}`);
				this.addresses = this.addresses.filter((addr) => addr.id !== id);
			},
		},
	};
</script>