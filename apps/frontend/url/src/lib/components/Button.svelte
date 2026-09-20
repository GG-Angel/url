<script lang="ts">
	import type { Snippet } from 'svelte';
	import type { HTMLButtonAttributes } from 'svelte/elements';

	type Props = HTMLButtonAttributes & {
		children: Snippet;
		variant?: 'primary' | 'secondary';
		class?: string;
		isIconOnly?: boolean;
	};

	let {
		children,
		variant = 'primary',
		isIconOnly = false,
		class: className,
		...props
	}: Props = $props();
</script>

<button
	{...props}
	class={[
		'flex cursor-pointer items-center gap-2 rounded font-medium transition-colors',
		isIconOnly ? 'p-2' : 'px-4 py-2',
		variant === 'primary'
			? 'bg-primary text-primary-foreground hover:bg-primary-hover'
			: 'bg-secondary text-secondary-foreground hover:bg-secondary-hover',
		className
	]}
>
	{@render children()}
</button>
