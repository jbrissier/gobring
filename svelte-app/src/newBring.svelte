<script>
    import {createEventDispatcher, onMount} from 'svelte'
    import {store, brings} from './store.js'
    import Time from './time.svelte'
    import _ from 'lodash'
    let where;
    let ref = null;
    let time = "";
    let defaultOptions = ["Aldi", "Lidl", "Rewe", "Edka", "Döner", "Pizza", "Curry", "sonstiges"]
    let selected = defaultOptions[0];




    const dispatch = createEventDispatcher();

    function saveNewBring() {

        fetch("/api/bring", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify({ where: where, until: _.toInteger( time), user: $store.username }) }).then(data => {

            dispatch("change", data)
            brings.update(brings => [...brings, data])


        });
    }

    function checkEnter(e){
        console.log(e)
        if(e.key === "Enter"){
            saveNewBring()
        }

    }



</script>

<div class="flex flex-col pt-1 w-1/2 align-middle justify-center">
    <p>Bringst du was mit {$store.username}?</p>


    <input class="mt-2 pl-2 p-4 mb-2" bind:this={ref} type="text" placeholder="Where" bind:value={where} on:keypress={checkEnter} id="where" />

    <Time bind:time={time} />
    <button on:click={saveNewBring} class="text-slate-100 p-4 mb-2 bg-slate-600 shadow-md ">  ok</button>
</div>
