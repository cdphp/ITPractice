<template>
	<section>
		<MyHeader></MyHeader>
		<router-view></router-view>
		<MyFooter></MyFooter>
	</section>
</template>

<script>
	import MyHeader from '../public/header.vue'
	import MyFooter from '../public/footer.vue'
	export default {
		data() {
			return {

			}
		},
		components: {
			MyHeader,
			MyFooter
		},
		mounted() {
			
		}
	}

</script>

<style scoped>

</style>
