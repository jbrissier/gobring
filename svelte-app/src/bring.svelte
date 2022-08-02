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

<div class="bring  bg-slate-900 text-slate-100 p-5 mt-1">
    
    <div class="info">
        <div class="where">{data.where} [{data.id}]</div>
        <div class="detail">
            <div>@{data.until}
             
        </div>
            <div>{data.user}</div>
        </div>
    </div>

    <div class="plus">+</div>

</div>

<button on:click={deleteBring}>delete</button>



