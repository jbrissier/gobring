<script>
    import { onMount } from 'svelte'
    import _ from "lodash";
    import moment from 'moment'
    
    let innterTime = moment().add(10, 'minutes') ;
    let saveInput = innterTime.format("HH:mm")
    export let time = innterTime ;

    onMount(()=>{
        time = innterTime;
    })

    function formatTime(e) {

        if(e.key == "Delete" || e.key == "Backspace"){
            return
        }
        let v = e.target.value;

        if (v.length === 2 && v[2] !== ":") {
            e.target.value = v + ":";
        }

        if(v.length === 6){
            e.target.value = saveInput
            return
        }
        if(e.keyCode == '38'){
            innterTime = innterTime.add(5, "minutes")
            e.target.value = innterTime.format("HH:mm")
        
        }
        if(e.keyCode == '40'){
            innterTime = innterTime.add(-5, "minutes")
            e.target.value = innterTime.format("HH:mm")
        
        }
        saveInput = e.target.value


        time = moment(saveInput, "HH:mm")
        innterTime = time;

   

    }

  

</script>

<input
    type="text"
    placeholder="Time"
    value={saveInput}
    class="pl-2 p-4 mb-2"

    on:keyup={formatTime}
    id="time"
/>
