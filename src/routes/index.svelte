<script>
	import { onMount } from "svelte";
	import Button from "../components/Button.svelte";

	let data;

	onMount(update);

	async function update() {
		data = null;
		const res = await fetch("/api/generate");
		data = await res.json();
	}
</script>

<div class="container">
	<div class="item" id="name">
		{#if data}
			{data.noun}<span>{data.name}</span>{data.number}
		{:else}
			...
		{/if}
	</div>
	<Button action={update} />
</div>

<style lang="scss">
	:global {
		body {
			background-color: rgb(53, 53, 53);
		}
	}

	@font-face {
		font-family: bignoodle;
		src: url('/big_noodle_titling_oblique.ttf');
		font-display: swap;
	}

	#name {
		font-family: bignoodle, sans-serif;
		font-size: max(7vw, 5.25rem);
		color: rgb(250, 156, 30);
	}

	span {
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
