<script>
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    let add_ref = null;
    let add = false;
    export let value = "";

    function showAdd() {
        add = true;
    }

    function keyUp(e) {
        if (e.key === "Enter") {
            add = false;
            dispatch("newBring", {value});
        }
    }

    $: {
        if (add && add_ref) {
            add_ref.focus();
        }
    }
</script>

<div class="plus cursor-pointer my-4" on:click={showAdd}>+</div>
{#if add}
    <input
        bind:this={add_ref}
        bind:value
        on:keyup={keyUp}
        type="text"
        placeholder="Item"
        class="description ml-4 mt-4 pl-2 p-4 mb-2"
    />
{/if}
