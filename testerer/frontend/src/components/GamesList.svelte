<script lang="ts">
    import { Grid } from "@svar-ui/svelte-grid";
    import { GetSaves } from '../../wailsjs/go/main/App';
    

    function doSmthn() {
        GetSaves()
        .then(response => response.data)
        .catch(error => error);
    }

    async function getResponse() {
        const response = await doSmthn();
        console.log(response)
    }
    
    getResponse()

    let data: Array<Record<string, any>> = GetSaves();

    const columns = [
        {id: "ID", header: "Unique Save Number"},
        {id: "Game_Name", header: "Game Name"},
        {id: "Console", header: "Console"},
        {id: "Device", header: "Current Device"},
        {id: "Time_Modified", header: "Last Saved"},
        {id: "Save_Path", header: "Save Location"},
    ]

    
</script>


<div class="h-full w-full">
    <div class="h-full grid grid-cols-3 grid-rows-6 gap-4">
        <div class="col-span-3 row-span-6">
          {data}
        </div>
    </div>
</div>

<style>


</style>
