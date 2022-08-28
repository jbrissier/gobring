<script>
    import { createEventDispatcher, onMount } from "svelte";
    import { store, brings } from "./store.js";
    import Time from "./time.svelte";
    import _ from "lodash";
import Brings from "./brings.svelte";
import BringItem from "./bringItem.svelte";
import Bring from "./bring.svelte";
    let where;
    let time = "";
    let show = false;
    let ref = null;

    const dispatch = createEventDispatcher();

    function saveNewBring() {
        fetch("/api/bring", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                where: where,
                until: _.toInteger(time),
                user: $store.username,
            }),
        }) 
            .then((data) => data.json())
            .then((data) => {
                dispatch("change", data);
                brings.update((b) => {return [...b, data]});
            });
        show = false;
    }

    function checkEnter(e) {
        if (e.key === "Enter") {
            saveNewBring();
        }
    }

    function showForm() {
        show = !show;
        if (ref) {
            ref.focus();
        }
    }
    $: {
        if (ref) {
            ref.focus();
        }
    }
</script>

{#if show}
    <div
        class="flex flex-col pt-1 w-11/12 md:w-1/2 align-middle justify-center absolute bg-white p-10 shadow-xl"
    >
        <div class="flex justify-between items-center">

            <h3 class="py-5 text-xl">Wo gehts hin, was bringst du mit?</h3>
            <div class="cursor-pointer p-3" on:click={()=>show=false}>x</div>
        </div>

        <input
            bind:this={ref}
            class="mt-2 pl-2 p-4 mb-2 outline-none"
            type="text"
            placeholder="Where"
            bind:value={where}
            on:keypress={checkEnter}
            id="where"
        />

        <Time bind:time />
        <button
            on:click={saveNewBring}
            class="text-slate-100 p-4 mb-2 bg-slate-600 shadow-md "
        >
            ok</button
        >
    </div>
{:else}
    <div on:click={showForm} class="cursor-pointer text-4xl mt-10">+</div>
{/if}