import networkx as nx
import matplotlib.pyplot as plt


def kruskal(graph):
    # Initialize the minimal island tree (MST) as an empty graph
    mst = nx.Graph()

    # Create a list of all edges sorted by weight
    edges = [(u, v, data['weight']) for u, v, data in graph.edges(data=True)]
    edges.sort(key=lambda x: x[2])

    # Create a disjoint-set data structure to track connected components
    disjoint_set = {node: node for node in graph.nodes()}

    def find_set(node):
        if disjoint_set[node] != node:
            disjoint_set[node] = find_set(disjoint_set[node])
        return disjoint_set[node]

    def union_sets(node1, node2):
        root1 = find_set(node1)
        root2 = find_set(node2)
        if root1 != root2:
            disjoint_set[root1] = root2

    for u, v, w in edges:
        if find_set(u) != find_set(v):
            mst.add_edge(u, v, weight=w)
            union_sets(u, v)

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

# Find the minimal island tree (MST) using Kruskal's algorithm
mst = kruskal(G)

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
