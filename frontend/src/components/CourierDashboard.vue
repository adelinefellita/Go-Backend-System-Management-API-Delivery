<template>
	<div>
		<h1>Courier Dashboard</h1>
		<ul>
			<li v-for="address in addresses" :key="address.id">
				{{ address }}
				<button @click="updateStatus(address.id)">Mark as Delivered</button>
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
			const response = await axios.get('http://localhost:8080/courier/addresses', {
				headers: {
					Authorization: `Bearer ${this.$store.state.token}`
				},
			});
			this.addresses = response.data;
		},
		methods: {
			async updateStatus(id) {
				await axios.put(`http://localhost:8080/courier/addresses/${id}/status`, {
					status: 'pengiriman selesai',
				});
				this.addresses = this.addresses.filter((addr) => addr.id !== id);
			},
		},
	};
</script>