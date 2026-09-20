<script lang="ts">
	import { House, Link, List, Tag } from '@lucide/svelte';
	import { page } from '$app/state';
	import NavItem from './NavItem.svelte';
</script>

<nav
	class="flex shrink-0 flex-col"
	aria-label="Sidebar"
>
	<NavItem
		path="/"
		label="Overview"
		icon={House}
		isActive={page.url.pathname === '/'}
	/>
	<NavItem
		path="/urls"
		label="List short URLs"
		icon={List}
		isActive={page.url.pathname.startsWith('/urls')}
	/>
	<NavItem
		path="/create"
		label="Create short URL"
		icon={Link}
		isActive={page.url.pathname.startsWith('/create')}
	/>
	<NavItem
		path="/tags"
		label="Manage tags"
		icon={Tag}
		isActive={page.url.pathname.startsWith('/tags')}
	/>
</nav>