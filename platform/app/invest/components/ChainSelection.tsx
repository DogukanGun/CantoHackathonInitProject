import { useEffect, useState } from "react";

export enum ChainTypes {
    Canto,
    Ethereum
}

type P = {
    chainSelection:(chain:string)=>void
}

const ChainSelection = ({chainSelection}:P) => {
    const [selectedOption, setSelectedOption] = useState<ChainTypes>(ChainTypes.Canto);

    useEffect(()=>{
        chainSelection(selectedOption == ChainTypes.Canto ? "canto" : "eth")
    },[selectedOption])

    return (
        <div className="flex items-center space-x-4">
            <div className="form-control">
                <label className="label cursor-pointer">
                    <span className="label-text text-black mr-2">Ethereum</span>
                    <input
                        type="radio"
                        value={ChainTypes.Ethereum}
                        name="radio-10"
                        className="radio checked:bg-red-500"
                        onChange={() => setSelectedOption(ChainTypes.Ethereum)}
                        checked={selectedOption === ChainTypes.Ethereum} />
                </label>
            </div>
            <div className="form-control">
                <label className="label cursor-pointer">
                    <span className="label-text text-black mr-2">Canto</span>
                    <input
                        type="radio"
                        name="radio-10"
                        value={ChainTypes.Canto}
                        className="radio checked:bg-blue-500"
                        checked={selectedOption === ChainTypes.Canto}
                        onChange={() => setSelectedOption(ChainTypes.Canto)}
                    />
                </label>
            </div>
        </div>
    )
}

export default ChainSelection;
