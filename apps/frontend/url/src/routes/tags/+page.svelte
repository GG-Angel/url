<script lang="ts">
	import { EllipsisVertical, Trash } from '@lucide/svelte';
	import type { PageProps } from './$types';
	import Button from '$lib/components/Button.svelte';
	import Dropdown from '$lib/components/Dropdown.svelte';
	import CardWithHeader from '$lib/components/CardWithHeader.svelte';
	import Card from '$lib/components/Card.svelte';

	let { data }: PageProps = $props();
</script>

<CardWithHeader header="Tags">
	<div class="table">
		<div class="table-header-group">
			<div class="table-row font-semibold">
				<div class="table-cell">Tag</div>
				<div class="table-cell">Links</div>
				<div class="table-cell"></div>
			</div>
		</div>
		<div class="table-row-group">
			{#each data.tags as tag (tag.id)}
				<div class="table-row">
					<div class="table-cell">{tag.name}</div>
					<div class="table-cell">{tag.total_links}</div>
					<div class="table-cell">
						<Dropdown placement="right">
							{#snippet clickable(triggerProps)}
								<Button {...triggerProps} variant="secondary" isIconOnly
									><EllipsisVertical /></Button
								>
							{/snippet}
							<Card class="p-1"
								><Button variant="danger"
									><Trash class="size-sm" /> Delete</Button
								>
							</Card>
						</Dropdown>
					</div>
				</div>
			{/each}
		</div>
	</div>
</CardWithHeader>
