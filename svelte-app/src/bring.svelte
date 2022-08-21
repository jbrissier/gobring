<script>
    import { onMount, createEventDispatcher } from "svelte";
    import {store} from './store.js' 
    import moment from 'moment'
    import BringItem from "./bringItem.svelte";
    import _ from "lodash";

    export let id = 0;
    export let data = {};
    let items = []
    let remaining = ""
    let diff = 0 

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
        data = newData;



    }

    function calculateRemaningTime(){

        let now = moment()
        diff = now.diff(moment(data.until))
        console.log(diff)
        remaining = moment.duration(diff).humanize()
    }
    onMount(() => {
  
        setInterval(calculateRemaningTime, 10000)


        if (!_.isEmpty(data)) {
            calculateRemaningTime()
            return;
        }

        fetchData().then((serverDate) => {
            data = serverDate;
            calculateRemaningTime()
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
    <div class="bring  bg-white shadow-xl text-slate-900 p-5 mt-1">


            <div class="flex justify-between where mb-5 align-center">
                <div>
                    <strong>{data.user}</strong> → <strong>{data.where}</strong> {moment(data.until).format("HH:mm")} ({remaining})
                </div>
                <div>
                    <button
                    class="text-xl "
                    on:click={deleteBring}>✕</button
                    >
                </div>
            </div>
         

        <div class="pl-3 mt-3">  

            {#if data.items}
            {#each data.items as item}
            
            <p>{item.user}: {item.description}</p>
            
            {/each}
            {/if}
        </div>
            
       
      

        <div class="flex items-center pl-3 align-middle">
            <BringItem on:newBring={newItem} />
        </div>


    </div>

</div>
