<script lang="ts">
	import type { Snippet } from 'svelte';

	type TriggerProps = {
		type: 'button';
		onclick: (event: MouseEvent) => void;
	};

	type Props = {
		clickable: Snippet<[TriggerProps]>;
		children: Snippet;
		placement?: 'left' | 'right';
	};

	let { clickable, children, placement = 'left' }: Props = $props();

	let isOpen = $state(false);

	const toggle = () => {
		isOpen = !isOpen;
	};

	const close = () => {
		isOpen = false;
	};

	const handleFocusOut = ({ relatedTarget, currentTarget }: FocusEvent) => {
		if (
			relatedTarget instanceof Node &&
			currentTarget instanceof Node &&
			currentTarget.contains(relatedTarget)
		)
			return;

		close();
	};

	const handleWindowKeydown = ({ key }: KeyboardEvent) => {
		if (isOpen && key === 'Escape') close();
	};
</script>

<svelte:window onkeydown={handleWindowKeydown} />

<div class="relative" onfocusout={handleFocusOut}>
	{@render clickable({
		type: 'button',
		onclick: () => toggle()
	})}
	{#if isOpen}
		<div
			class={[
				'absolute z-10 mt-2',
				placement === 'left' ? 'left-0' : 'right-0'
			]}
		>
			{@render children()}
		</div>
	{/if}
</div>
