<script lang="ts">
	import type { Snippet } from 'svelte';

	type Props = {
		children: Snippet;
		class?: string;
	};

	let { children, class: className }: Props = $props();
</script>

<section class={['flex flex-col rounded border bg-surface', className]}>
	{@render children()}
</section>
