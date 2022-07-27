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

    $: {
        where = selected
        if(selected === "sonstiges"){
            if(ref){
                ref.focus(); 

            }

        }
    }


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

<div>
    <p>Bringst du was mit {$store.username}?</p>

    {#if selected != "sonstiges"}
    <select bind:value={selected}>
        {#each defaultOptions as df}
            <option>{df}</option>
        {/each}
    </select>
    {/if}
    {#if selected === "sonstiges"}
        <input  bind:this={ref} type="text" placeholder="Where" bind:value={where} on:keypress={checkEnter} id="where" />
    {/if}
    <Time bind:time={time} />
    <button on:click={saveNewBring}> ok</button>
</div>
