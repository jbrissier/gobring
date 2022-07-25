<script>
    import {createEventDispatcher, onMount} from 'svelte'
    import Time from './time.svelte'
    import _ from 'lodash'
    let where;
    let ref = null;
    let time = "";


    onMount(() => {
        ref.focus(); 
    });


    const dispatch = createEventDispatcher();
    
    function saveNewBring() {
        fetch("/api/bring", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            body: JSON.stringify({ where: where, until: _.toInteger( time) }) }).then(data => {
      
            where = ""
            dispatch("change", data)
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
    <p>Bringst du was mit?</p>

    
    <input  bind:this={ref} type="text" placeholder="Where" bind:value={where} on:keypress={checkEnter} id="where" />
    <Time bind:time={time} />
    <button on:click={saveNewBring}> ok</button>
    {time}
</div>
