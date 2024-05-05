"use client"

import { useState } from "react";
import CustomButton from "@/components/Button";
import Dropdown from "@/components/Dropdown";

type Step2Params = {
  assetChosen: (asset: string) => void
}

const Step2 = ({ assetChosen }: Step2Params) => {
  const [asset, setAsset] = useState<string>("Lending".toLocaleLowerCase());

  const handleSubmit = async () => {
    assetChosen(asset.toLocaleLowerCase())
  }

  return (
    <div className="max-w-lg mx-auto">
      <h1 className="text-black my-2 capitalize">Which kind of asset you are holding ?</h1>
      <Dropdown setSelection={setAsset} options={["Lending","Dex"]} text="Invesment Type" />
      <CustomButton onClick={handleSubmit} text="Get Result" />
    </div>
  );
}

export default Step2;