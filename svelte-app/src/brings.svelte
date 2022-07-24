
<script>
    import Bring from "./bring.svelte";
    import { onMount } from "svelte";

    let brings = [];
    


    function fetchData() {
        return fetch(`http://localhost:8000/api/bring`)
            .then(response => response.json())
            .then(data => {
               
                brings = data;
            });
    }

    onMount(fetchData)


</script>
<h2>All Brings of the day</h2>
{#each brings as bring}

    <Bring id={bring.ID} data={bring}/>
{/each}