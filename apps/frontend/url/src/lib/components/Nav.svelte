<script lang="ts">
	import { House, Link, List, Tag } from '@lucide/svelte';
	import { page } from '$app/state';
	import NavItem from './NavItem.svelte';
</script>

<nav
	class="sidebar-nav flex shrink-0 flex-col"
	aria-label="Sidebar"
>
	<NavItem
		path="/"
		label="Overview"
		icon={House}
		active={page.url.pathname === '/'}
	/>
	<NavItem
		path="/urls"
		label="List short URLs"
		icon={List}
		active={page.url.pathname.startsWith('/urls')}
	/>
	<NavItem
		path="/create"
		label="Create short URL"
		icon={Link}
		active={page.url.pathname.startsWith('/create')}
	/>
	<NavItem
		path="/tags"
		label="Manage tags"
		icon={Tag}
		active={page.url.pathname.startsWith('/tags')}
	/>
</nav>