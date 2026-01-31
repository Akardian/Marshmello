import { Component, createSignal, Show } from "solid-js";
import "./Sidebar.css"


export const Sidebar: Component = () => {
	let [isOpen, setIsOpen] = createSignal(false);

	const toggle = (() => {
		setIsOpen(!isOpen())
	})

	return (
		<>
			<div classList={{
				sidebar: true,
				open: isOpen(),
				closed: !isOpen(),
			}}>
				<Show
					when={isOpen()}
					fallback={<button onClick={() => toggle()}>N</button>}
				>
					<button onClick={() => toggle()}>Navigation</button>
				</Show>


				<h2>Open: {true === isOpen() ? "true" : "false"}</h2>
			</div >
		</>
	);
};
