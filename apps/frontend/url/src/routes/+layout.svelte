<script lang="ts">
	import {  House, Link, List, Tag } from '@lucide/svelte';
	import { resolve } from '$app/paths';
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<div class="flex h-svh flex-col">
	<header class="border-b bg-background">
		<div class="flex items-center justify-between p-4 sm:px-6">
			<a class="font-mono text-sm font-semibold text-foreground" href={resolve('/')}>shrt</a>
		</div>
	</header>

	<div class="flex flex-1 min-h-0">
		<nav class="sidebar-nav w-54 flex shrink-0 flex-col border-r pt-4" aria-label="Sidebar">
			<a href={resolve('/')}><House /> Overview</a>
			<a href={resolve('/')}><List /> List short URLs</a>
			<a href={resolve('/')}><Link /> Create short URL</a>
			<a href={resolve('/')}><Tag /> Manage tags</a>
		</nav>
		<main
			class="flex flex-1 flex-col gap-8 overflow-y-auto p-4 sm:p-6"
		>
			{@render children()}
		</main>
	</div>
</div>
