<template>
	<section>
		<router-view></router-view>
		<MyFooter></MyFooter>
	</section>
</template>

<script>
	import MyHeader from '../public/header.vue'
	import MyFooter from '../public/footer.vue'
	export default {
		components: {
			MyHeader,
			MyFooter
		}
	}

</script>

<style scoped>

</style>
