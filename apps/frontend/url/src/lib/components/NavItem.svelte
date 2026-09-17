<script lang="ts">
	import type { Component } from 'svelte';
	import type { ResolvedPathname } from '$app/types';

	type Props = {
		path: ResolvedPathname;
		label: string;
		icon: Component;
		active?: boolean;
		iconOnly?: boolean;
	};

	let { path, label, icon: Icon, active = false, iconOnly = false }: Props = $props();
</script>

<a href={path} class:active>
	<Icon />
	{#if !iconOnly}
		{label}
	{/if}
</a>