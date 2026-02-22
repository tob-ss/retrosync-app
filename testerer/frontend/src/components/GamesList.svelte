<script lang="ts">
    import { GetSaves } from "../../wailsjs/go/main/App";
    import { GetHeaders } from "../../wailsjs/go/main/App";

  //  let data: Array<Record<string, any>> = GetSaves();

    GetSaves().then(data => {
        console.log(data);
    });

    let data = $state([]);

    (async () => {
        const data = await GetSaves();
        
    });

    GetHeaders(data).then(headers => {
        console.log(headers);
    });

    let headers = $state([]);

    (async () => {
        const headers = await GetHeaders(data);
    });

    
    
</script>


<div class="h-full w-full">
    <div class="h-full grid grid-cols-3 grid-rows-6 gap-4 text-[#D7D6FC] font-heebo">
        <div class="table col-span-3 row-span-6">
        {#await GetSaves()}
        <p>Getting Saves...</p>
        {:then data}
            {#await GetHeaders(data)}
            <p>Getting Headers...</p>
            {:then headers}
            {#each headers as header}
                <h1>{header}</h1>
                {#each data as save}
                <p>Showing game: {save.Game_Name}</p>
                {/each}
            {/each}
            {/await}
        {/await}
        </div>
    </div>
</div>