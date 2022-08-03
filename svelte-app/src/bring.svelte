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
<div class="flex flex-col w-1/2">

    <div class="bring  bg-slate-300 text-slate-900 p-5 mt-1">

        <div class="info">
            <div class="where">{data.user} geht zu {data.where} [{data.id}]</div>
            <div class="detail">
                <div>Bestellungen bis {' '} @{data.until}</div>

            </div>
    </div>

    <div class="plus">+</div>

</div>

<button class="text-slate-100 p-2 mb-2 bg-slate-600 shadow-md mt-1 self-end" on:click={deleteBring}>delete</button>

</div>


