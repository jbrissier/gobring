<script>
    import { onMount, createEventDispatcher } from "svelte";
    import {store} from './store.js' 
    import BringItem from "./bringItem.svelte";
    import _ from "lodash";

    export let id = 0;
    export let data = {};
    let items = []

    const dispatch = createEventDispatcher();

    function deleteBring() {
        fetch(`/api/bring/${data.id}`, {
            method: "DELETE",
        })
            .then((response) => response.json())
            .then((data) => {
                dispatch("change", data);
                console.log(data);
            });
    }

    function fetchData() {
        return fetch(`api/bring/${id}`)
            .then((response) => response.json())
            .then((data) => {
                console.log(data);
                return data;
            });
    }

    async function  updateBring(){

        const newData = await fetchData()
        console.log("......",newData)
        data = newData;



    }



    onMount(() => {
        console.log("data is not empty", data);
  
        if (!_.isEmpty(data)) {
            return;
        }

        fetchData().then((serverDate) => {
            data = serverDate;
        });
    });

    function fetchItems(){
        fetch(`api/bring/${id}/items`).then(data=>data.json()).then(jd=>{

            items = jd
            console.log(jd)


        })

    }



    function newItem(item) {
        
        const newItem = {
            user: $store.username,
            description: item.detail.value 
        }

        fetch(`api/bring/${id}`, { 
            method: "POST", 
            body: JSON.stringify(newItem) }
        ).then(res=>updateBring());
    }
</script>

<div class="flex flex-col w-1/2">
    <div class="bring  bg-slate-300 text-slate-900 p-5 mt-1">
        <div class="info">
            <div class="where">
                {data.user} geht zu {data.where} [{data.id}]
            </div>
            <div class="detail">
                <div>Bestellungen bis {" "} @{data.until}</div>
            </div>
        </div>

        <div class="pl-3 mt-3">  


            {#each data.items as item}
            
            <p>{item.user}: {item.description}</p>
            
            {/each}
        </div>
            
       
      

        <div class="flex items-center">
            <BringItem on:newBring={newItem} />
        </div>
    </div>

    <button
        class="text-slate-100 p-2 mb-2 bg-slate-600 shadow-md mt-1 self-end"
        on:click={deleteBring}>delete</button
    >
</div>
