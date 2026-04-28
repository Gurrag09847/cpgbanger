<script lang="ts">
	export const prerender = true;
	import Navbar from '@/components/navbar.svelte';
	import '../app.css';

	let { children } = $props();
</script>

<Navbar />
<main class="container mx-auto mt-12 mb-12 pt-12 pb-12">
	{@render children()}

	<!-- <img src="/pen.svg" class="w-sm" /> -->
</main>
<div class="container mx-auto mb-12 flex justify-center pb-12">
	<img src="/logo.svg" class="w-sm" />
</div>
