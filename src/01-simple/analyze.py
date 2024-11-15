import matplotlib.pyplot as plt
import pandas as pd

log_file = "action.log"

timestamps = []
philosophers = []
actions = []

with open(log_file, mode='r') as f:
    for line in f:
        if "Dinner is over." in line:
            continue
        parts = line.strip().split(" ")
        time = parts[0][1:-1] # timestamps
        philosopher = int(parts[2])  # philosophers
        action = " ".join(parts[4:])  # actions

        timestamps.append(pd.to_datetime(time, format="%H:%M:%S.%f"))
        philosophers.append(philosopher)
        actions.append(action)

data = pd.DataFrame({
    "Timestamp": timestamps,
    "Philosopher": philosophers,
    "Action": actions
})

action_colors = {
    "thinking (wait for forks)": "#812990",
    "eating": "#5eb954",
    "sleeping": "#f19db5",
    "finished eating": "#7cc7e8"
}

data["Color"] = data["Action"].map(action_colors)
data = data[data["Color"].notnull()]

plt.figure(figsize=(14, 6))

for philosopher in data["Philosopher"].unique():
    subset = data[data["Philosopher"] == philosopher]
    for i in range(len(subset) - 1):
        start_time = subset.iloc[i]["Timestamp"]
        end_time = subset.iloc[i + 1]["Timestamp"]
        action = subset.iloc[i]["Action"]
        color = subset.iloc[i]["Color"]

        if color is None or not isinstance(color, str):
            raise ValueError(f"Invalid color for action '{action}' at index {i}")

        if subset.iloc[i]["Action"] == "finished eating":
            plt.plot(start_time, philosopher, 'o', color=subset.iloc[i]["Color"], markersize=10)
        else:
            plt.hlines(
                y=philosopher,
                xmin=start_time,
                xmax=end_time,
                colors=subset.iloc[i]["Color"],
                linewidth=6,
                label=action if i == 0 else None
            )
    # last "finished eating"
    last_row = subset.iloc[-1]
    if last_row["Action"] == "finished eating":
        plt.plot(last_row["Timestamp"], philosopher, 'o', color=last_row["Color"], markersize=10)

legend_handles = [plt.Line2D([0], [0], color=color, lw=6, label=action) for action, color in action_colors.items()]
plt.legend(handles=legend_handles, title="Actions", loc="upper left", bbox_to_anchor=(1, 1))


plt.title("Dining Philosophers Timeline")
plt.xlabel("Time")
plt.ylabel("Philosophers")
plt.yticks(data["Philosopher"].unique())
plt.grid(True)
plt.tight_layout()
plt.show()