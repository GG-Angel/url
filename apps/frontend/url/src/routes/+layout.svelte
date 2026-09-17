<script lang="ts">
	import { resolve } from '$app/paths';
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import Nav from '$lib/components/Nav.svelte';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="flex h-svh flex-col">
	<header class="border-b bg-background">
		<div class="flex items-center justify-between p-4 sm:px-6">
			<a
				class="font-mono text-sm font-semibold text-foreground"
				href={resolve('/')}>shrt</a
			>
		</div>
	</header>

	<div class="flex flex-col-reverse min-h-0 flex-1 sm:flex-row">
		<div class="flex flex-col justify-between w-auto border-t sm:border-t-0 sm:border-r py-4 sm:w-56">
			<Nav />
		</div>
		<main class="flex flex-1 flex-col gap-8 overflow-y-auto p-4 sm:p-6">
			{@render children()}
		</main>
	</div>
</div>
