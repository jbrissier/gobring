<script>
    import Bring from "./bring.svelte";
    import NewBring from "./newBring.svelte";
    import { onMount } from "svelte";

    let brings = [];
    


    function fetchData() {
        return fetch(`/api/bring`)
            .then(response => response.json())
            .then(data => {
               
                brings = data;
            });
    }

    onMount(fetchData)


</script>
<h2>All Brings of the day</h2>
<NewBring on:change={fetchData}/>

{#each brings as bring}

    <Bring on:change={fetchData} id={bring.ID} data={bring}/>
{/each}
<hr/>