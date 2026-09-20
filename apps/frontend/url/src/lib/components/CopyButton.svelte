<script lang="ts">
	import { Check, Copy } from '@lucide/svelte';
	import Button from './Button.svelte';

	type Props = {
		copyText: string;
	};

	let { copyText }: Props = $props();
	let isCopyActive = $state(false);

	const handleClick = async () => {
		await navigator.clipboard.writeText(copyText);

		isCopyActive = true;

		setTimeout(() => {
			isCopyActive = false;
		}, 2000);
	};
</script>

<Button variant="secondary" isIconOnly={true} onclick={handleClick}>
	{#if isCopyActive}
		<Check class="size-sm text-success" />
	{:else}
		<Copy class="size-sm" />
	{/if}
</Button>
