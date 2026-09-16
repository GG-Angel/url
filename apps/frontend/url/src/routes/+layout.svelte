<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<header class="bg-background border-b">
	<div class="mx-auto flex h-14 max-w-5xl items-center justify-between px-4 sm:px-6">
		<a class="text-foreground font-mono text-sm font-semibold" href="/">shrt</a>
	</div>
</header>
<main class="mx-auto flex w-full max-w-5xl flex-col gap-8 px-4 py-8 sm:px-6 sm:py-12">
	{@render children()}
</main>

