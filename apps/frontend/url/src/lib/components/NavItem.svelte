<script lang="ts">
	import type { Component } from 'svelte';
	import type { ResolvedPathname } from '$app/types';

	type Props = {
		path: ResolvedPathname;
		label: string;
		icon: Component;
		isActive?: boolean;
		isIconOnly?: boolean;
	};

	let {
		path,
		label,
		icon: Icon,
		isActive = false,
		isIconOnly = false
	}: Props = $props();
</script>

<a
	href={path}
	class={[
		'flex items-center gap-2.5 bg-background px-4 py-2',
		isActive
			? 'bg-primary text-primary-foreground font-medium'
			: 'text-foreground hover:bg-secondary-hover hover:text-foreground'
	]}
>
	<Icon class="size-sm" />
	{#if !isIconOnly}
		{label}
	{/if}
</a>
