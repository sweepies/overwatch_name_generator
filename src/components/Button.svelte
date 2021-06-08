<script>
    export let action;

    let spinning;

    let spin = async () => {
        spinning = true;
    };

    let unspin = async () => {
        spinning = false
    };
</script>

<button on:click={action} on:click={spin} on:transitionend={unspin} class:rotate={spinning}>
    <i class="fas fa-redo"></i>
</button>

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

    button.rotate {
        transition: all 0.5s;
        transform: rotate(360deg);
    }
</style>
