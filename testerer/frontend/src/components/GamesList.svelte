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

    let headers = $state([]);

    function RetrieveHeaders(data): string[] {
        GetHeaders(data).then(headers => {
            console.log(headers);
        });

        (async () => {
            const headers = await GetHeaders(data);
        });
        return headers
    } 

    

   
    
    
</script>


<div class="h-full w-full">
    <div class="h-full grid grid-cols-3 grid-rows-6 gap-4 text-[#D7D6FC] font-heebo">
        {#await GetSaves()}
            {$inspect(data)}
        <p>Getting Saves...</p>
        {:then data}
            {#await RetrieveHeaders(data)}
            {$inspect(data)}
            <p>Getting Headers...</p>
            {:then headers}
            {$inspect(data)}
            {$inspect(headers)}
                {#each headers as header}
                {$inspect(data)}
                {$inspect(headers)}
                    <h1>{header}</h1>
                    {#each data as save}
                    <p>Showing game: {save.Game_Name}</p>
                    {/each}
                {/each}
            {/await}
        {/await}
    </div>
</div>