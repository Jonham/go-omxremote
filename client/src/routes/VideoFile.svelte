<script>
	import Button from './Button.svelte';

	export let file;
	export let hash;

	let activecontrols = false;

	const clickVideo = (e) => {
		activecontrols = !activecontrols;
	};
</script>

<div class="w-full rounded overflow-hidden shadow p-2">
	<header class="p-2" on:click={clickVideo}>
		<p class="font-semibold">{file}</p>
	</header>
	{#if activecontrols === true}
		<div class="flex-row my-1">
			<Button action="start" text="Start" {hash} />
			<Button action="pause" text="Pause" />
			<Button action="ShowInfo" text="ShowInfo" />
		</div>
		<div class="flex-row my-1">
			<Button action="backward" text="Back" />
			<Button action="forward" text="Forward" />
			<Button action="stop" text="Stop" />
		</div>
		<div class="flex-row my-1">
			<Button action="PreAudio" text="PreAudio" />
			<Button action="NextAudio" text="NextAudio" />
		</div>
		<div class="flex-row my-1">
			<Button action="IncreaseVolume" text="IncreaseVolume" />
			<Button action="DecreaseVolume" text="DecreaseVolume" />
		</div>
		<div class="flex-row my-1">
			<Button action="prevsubs" text="Prev subs" />
			<Button action="nextsubs" text="Next subs" />
			<Button action="ToggleSubtitle" text="ToggleSubtitle" />
		</div>
	{/if}
</div>

<style>
	header {
		cursor: pointer;
	}
	.flex-row {
		margin-bottom: 5px;
	}
</style>
