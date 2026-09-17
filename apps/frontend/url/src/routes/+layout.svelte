<script lang="ts">
	import { House, Link, List, Tag } from '@lucide/svelte';
	import { resolve } from '$app/paths';
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import { page } from '$app/state';

	let { children } = $props();

	const navs = [
		{
			path: '/',
			label: 'Overview',
			Icon: House,
			isActive: (path: string) => path === '/'
		},
		{
			path: '/urls',
			label: 'List short URLs',
			Icon: List,
			isActive: (path: string) => path.startsWith('/urls')
		},
		{
			path: '/create',
			label: 'Create short URL',
			Icon: Link,
			isActive: (path: string) => path.startsWith('/create')
		},
		{
			path: '/tags',
			label: 'Manage tags',
			Icon: Tag,
			isActive: (path: string) => path.startsWith('/tags')
		}
	] as const;
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

	<div class="flex min-h-0 flex-1">
		<nav
			class="sidebar-nav flex w-54 shrink-0 flex-col border-r pt-4"
			aria-label="Sidebar"
		>
			{#each navs as { path, label, Icon, isActive } (path)}
				<a
					href={resolve(path)}
					class={isActive(page.url.pathname) ? 'active' : ''}
				>
					<Icon />
					{label}
				</a>
			{/each}
		</nav>
		<main class="flex flex-1 flex-col gap-8 overflow-y-auto p-4 sm:p-6">
			{@render children()}
		</main>
	</div>
</div>
