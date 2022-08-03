<script>
    import Bring from "./bring.svelte";
    import NewBring from "./newBring.svelte";
    import {brings} from './store.js'
    import { onMount } from "svelte";






    function fetchData() {
        return fetch(`/api/bring`)
            .then(response => response.json())
            .then(data => {



                brings.set(data)

            });
    }

    onMount(fetchData)


</script>




{#each $brings as bring}

<Bring on:change={fetchData} id={bring.ID} data={bring}/>
{/each}

<NewBring/>