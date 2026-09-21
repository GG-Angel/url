<script lang="ts">
	import Card from "$lib/components/Card.svelte";
	import type { PageProps } from "./$types";

    let { data }: PageProps = $props();
</script>

<Card header="Tags">
    <div class="table">
        <div class="table-header-group">
            <div class="table-row font-semibold">
                <div class="table-cell">Tag</div>
                <div class="table-cell">Links</div>
            </div>
        </div>
        <div class="table-row-group">
            {#each data.tags as tag (tag.id)}
                <div class="table-row">
                    <div class="table-cell">{tag.name}</div>
                    <div class="table-cell">{tag.total_links}</div>
                </div>
            {/each}
        </div>
    </div>
</Card>