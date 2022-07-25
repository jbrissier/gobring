<script>
import { onMount, createEventDispatcher } from 'svelte';
import _ from 'lodash'


    export let id = 0;
    export let data = {}
    const dispatch = createEventDispatcher();

    function deleteBring(){
        fetch(`/api/bring/${data.id}`, {
            method: 'DELETE'
          
        })
        .then(response => response.json())
        .then(data => {
            dispatch("change", data)
            console.log(data)
        })
    }

    function fetchData() {
        return fetch(`api/bring/${id}`)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                return data;
            });
    }

    onMount(() => {
        
        console.log('data is not empty', data)
        if(!_.isEmpty(data)) {
            return
        }

        fetchData().then(serverDate => {
            data = serverDate;
        });
    });

</script>

<hr/>
<h2>Bring {data.where}  [{data.id}]</h2>
<p>{data.description}</p>
<p>@{data.until}</p>

<button on:click={deleteBring}>Delete</button>