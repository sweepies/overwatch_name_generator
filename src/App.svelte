<script>
	import { onMount } from "svelte";
	import Button from "./Button.svelte";

	let data;

	onMount(update);

	async function update() {
		data = null;
		const res = await fetch("/api/generate");
		data = await res.json();
	}
</script>

<div class="container">
	{#if data}
		<div class="item" id="name">
			{data.noun}<span id="name-name">{data.name}</span>{data.number}
		</div>
	{:else}
		<div class="item" id="name">...</div>
	{/if}
	<Button action={update} />
</div>

<style>
	#name {
		font-family: bignoodle, sans-serif;
		font-size: max(7vw, 5.25rem);
		color: rgb(250, 156, 30);
	}

	#name-name {
		color: rgb(180, 180, 180);
	}

	.container {
		display: flex;
		align-items: center;
		justify-content: center;
		margin-top: 30vh;
		flex-direction: column;
		align-content: space-between;
	}

	.item {
		display: flex;
		margin-bottom: 5vh;
	}
</style>
