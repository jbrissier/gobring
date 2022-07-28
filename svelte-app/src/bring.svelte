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

<div class="bring">
    
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



<style>

    .bring{
        margin: 1rem 0;
        padding: 3rem;
        border-left: 1px solid #eee;
        display: flex;
        flex-direction: column;
    }

    .plus{
		font-size: 2rem;
		cursor: pointer;
	}

    .bring .info{
        display: flex;
 
        align-items: center;
        
    }
    .where{
        margin-right: 2rem;
        font-size: 1.5rem;
        
    }
    .bring .info .detail{
        display: flex;
        flex-direction: column;
    }

</style>
