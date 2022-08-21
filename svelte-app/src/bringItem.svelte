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
            value = ""
        }
    }

    $: {
        if (add && add_ref) {
            add_ref.focus();
        }
    }
</script>

<div class="plus cursor-pointer  text-2xl flex align-center" on:click={showAdd}>+</div>
{#if add}
    <input
        bind:this={add_ref}
        bind:value
        on:keyup={keyUp}
        type="text"
        placeholder="Ich brauche..."
        class="description pl-2 p-4  outline-none"
    />
{/if}
