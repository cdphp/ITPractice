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
				test:123,
			}
		},
		components: {
			MyHeader,
			MyFooter
		},
		mounted() {
			console.log("mounted from main:",this.test)
		}
	}

</script>

<style scoped>

</style>
