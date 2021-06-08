<script>
    export let action;

    import { Motion } from 'svelte-motion';
    let i = 0;
</script>

<Motion animate={{ rotate: i }} transition={{ duration: .5 }} let:motion>
    <button aria-label="refresh" on:click={action} on:click={ () => i+=360 } use:motion>
        <i class="fas fa-redo"></i>
    </button>
</Motion>

<style>
    button {
        background-color: rgb(250, 156, 30);
        border: none;
        color: rgb(53, 53, 53);
        font-size: max(2vw, 1.5rem);
        padding: 1rem;
        border-radius: 100%;
        cursor: pointer;
        display: flex;
    }

    button:hover {
        background-color: rgb(192, 122, 31);
    }
</style>
