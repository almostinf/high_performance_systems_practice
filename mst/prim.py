import networkx as nx
import matplotlib.pyplot as plt


def prim(graph):
    # Create an empty set to store selected vertices
    selected = set()

    # Select an arbitrary starting node
    start_node = list(graph.nodes())[0]
    selected.add(start_node)

    # Initialize the minimum spanning tree (MST) as an empty graph
    mst = nx.Graph()

    while len(selected) < len(graph.nodes()):
        edge_candidates = []
        for node in selected:
            for neighbor, edge_data in graph[node].items():
                if neighbor not in selected:
                    edge_candidates.append((node, neighbor, edge_data['weight']))
        # Check if there are any edge candidates left
        if not edge_candidates:
            break
        # Choose the edge with the smallest weight
        edge_candidates.sort(key=lambda x: x[2])
        u, v, w = edge_candidates[0]
        mst.add_edge(u, v, weight=w)
        selected.add(v)

    return mst

# wiki example
G = nx.Graph()
G.add_edge('A', 'B', weight=7)
G.add_edge('A', 'D', weight=5)
G.add_edge('B', 'C', weight=8)
G.add_edge('B', 'E', weight=7)
G.add_edge('C', 'E', weight=5)
G.add_edge('D', 'E', weight=15)
G.add_edge('D', 'F', weight=6)
G.add_edge('F', 'G', weight=11)
G.add_edge('G', 'E', weight=9)
G.add_edge('F', 'E', weight=8)
G.add_edge('B', 'D', weight=9)

# Check if the graph is connected
if not nx.is_connected(G):
    components = list(nx.connected_components(G))
    msts = []
    for component in components:
        subgraph = G.subgraph(component)
        mst = prim(subgraph)
        msts.append(mst)
    mst = nx.compose_all(msts)
else:
    mst = prim(G)

# Visualize the original graph and the MST
pos = nx.spring_layout(G)
plt.figure(figsize=(12, 5))

plt.subplot(121)
nx.draw(G, pos, with_labels=True, node_color='lightblue', node_size=500, font_size=10, font_color='black')
nx.draw_networkx_edge_labels(G, pos, edge_labels={(u, v): G[u][v]['weight'] for (u, v) in G.edges()})
plt.title("Original Graph")

plt.subplot(122)
nx.draw(mst, pos, with_labels=True, node_color='lightblue', node_size=500, font_size=10, font_color='black')
nx.draw_networkx_edge_labels(mst, pos, edge_labels={(u, v): mst[u][v]['weight'] for (u, v) in mst.edges()})
plt.title("Minimum Spanning Tree (MST)")

plt.tight_layout()
plt.show()
