<script lang="ts">
    import { Grid } from "@svar-ui/svelte-grid";
    import { GetSaves } from '../../wailsjs/go/main/App';
    

  //  let data: Array<Record<string, any>> = GetSaves();

    let loadGrid: boolean = $state(false);

    const columns = [
            {id: "ID", header: "Unique Save Number", flexgrow: 0.25, sort: true},
            {id: "Game_Name", header: "Game Name", flexgrow: 1, sort: true},
            {id: "Console", header: "Console", flexgrow: 1, sort: true},
            {id: "Device", header: "Current Device", flexgrow: 0.5, sort: true},
            {id: "Time_Modified", header: "Last Saved", flexgrow: 0.5, sort: true},
            {id: "Save_Path", header: "Save Location", flexgrow: 2, sort: true},
        ]

    GetSaves().then(data => {
        console.log(data);
    });

    let data = $state([]);

    (async () => {
        const data = await GetSaves();
    });
    
</script>


<div class="h-full w-full">
    <div class="h-full grid grid-cols-3 grid-rows-6 gap-4">
        <div class="col-span-3 row-span-6">
        {#await GetSaves()}
            <p>Getting Saves...</p>
        {:then data}
            <Grid {data} {columns} />
        {/await}
        </div>
    </div>
</div>

<style>


</style>
